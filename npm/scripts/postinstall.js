#!/usr/bin/env node

/**
 * Postinstall script for @pwarnock/kjbc-mcp
 * Downloads the correct platform-specific Go binary from GitHub releases
 */

const fs = require("fs");
const path = require("path");
const https = require("https");
const { execFileSync } = require("child_process");

const pkg = require("../package.json");
const REPO = "pwarnock/kjbc-mcp";
const BINARY_NAME = "kjbc-mcp";

function log(msg) {
  console.log(`[${pkg.name}] ${msg}`);
}

function error(msg) {
  console.error(`[${pkg.name}] ${msg}`);
}

/**
 * Detect platform and architecture
 */
function getPlatformInfo() {
  const platform = process.platform;
  const arch = process.arch;

  // Map Node.js platform/arch to Go naming convention
  const platformMap = {
    darwin: "darwin",
    linux: "linux",
    win32: "windows",
  };

  const archMap = {
    x64: "amd64",
    arm64: "arm64",
  };

  const goPlatform = platformMap[platform];
  const goArch = archMap[arch];

  if (!goPlatform || !goArch) {
    return null;
  }

  return {
    platform: goPlatform,
    arch: goArch,
    extension: platform === "win32" ? ".exe" : "",
  };
}

/**
 * Download file with redirect support
 */
function download(url) {
  return new Promise((resolve, reject) => {
    const request = https.get(url, (response) => {
      // Handle redirects (GitHub releases use them)
      if (response.statusCode >= 300 && response.statusCode < 400 && response.headers.location) {
        return download(response.headers.location).then(resolve).catch(reject);
      }

      if (response.statusCode !== 200) {
        return reject(new Error(`HTTP ${response.statusCode}: ${response.statusMessage}`));
      }

      const chunks = [];
      response.on("data", (chunk) => chunks.push(chunk));
      response.on("end", () => resolve(Buffer.concat(chunks)));
      response.on("error", reject);
    });

    request.on("error", reject);
    request.setTimeout(60000, () => {
      request.destroy();
      reject(new Error("Download timeout"));
    });
  });
}

/**
 * Extract tar.gz using system tar (available on macOS, Linux, and modern Windows)
 * Uses execFileSync to avoid shell injection
 */
function extractTarGz(tarBuffer, destDir) {
  const tarPath = path.join(destDir, "temp.tar.gz");
  fs.writeFileSync(tarPath, tarBuffer);

  try {
    // Use execFileSync instead of execSync to avoid shell injection
    execFileSync("tar", ["-xzf", tarPath, "-C", destDir], { stdio: "pipe" });
  } finally {
    fs.unlinkSync(tarPath);
  }
}

/**
 * Main installation
 */
async function main() {
  const platformInfo = getPlatformInfo();

  if (!platformInfo) {
    error(`Unsupported platform: ${process.platform}/${process.arch}`);
    error("Please download the binary manually from:");
    error(`https://github.com/${REPO}/releases`);
    process.exit(0); // Don't fail the install
  }

  const { platform, arch, extension } = platformInfo;
  const version = pkg.version;
  const archiveName = `${BINARY_NAME}_${platform}_${arch}.tar.gz`;
  const url = `https://github.com/${REPO}/releases/download/v${version}/${archiveName}`;

  log(`Downloading ${BINARY_NAME} v${version} for ${platform}/${arch}...`);

  try {
    const tarBuffer = await download(url);
    log("Extracting...");

    const binDir = path.join(__dirname, "..", "bin");
    extractTarGz(tarBuffer, binDir);

    // Ensure binary is executable
    const binaryPath = path.join(binDir, BINARY_NAME + extension);
    if (process.platform !== "win32") {
      fs.chmodSync(binaryPath, 0o755);
    }

    log("Installation complete!");
  } catch (err) {
    error(`Download failed: ${err.message}`);
    error("");
    error("You can install the binary manually:");
    error(`1. Download from: https://github.com/${REPO}/releases`);
    error(`2. Extract to: ${path.join(__dirname, "..", "bin")}`);
    process.exit(0); // Don't fail the install - allows manual setup
  }
}

main();
