const os = require('os');
const fs = require('fs');
const path = require('path');

const platform = os.platform();
let binaryPath = '';
const destPath = path.join(__dirname, 'spix');

if (platform === 'darwin') {
  binaryPath = path.join(__dirname, 'dist', 'spix-macos');
} else if (platform === 'linux') {
  binaryPath = path.join(__dirname, 'dist', 'spix-linux');
} else if (platform === 'win32') {
  binaryPath = path.join(__dirname, 'dist', 'spix.exe');
} else {
  console.error(`Unsupported platform: ${platform}`);
  process.exit(1);
}

fs.copyFileSync(binaryPath, destPath);

if (platform === 'win32') {
  fs.renameSync(destPath, destPath + '.exe');
}

console.log(`spix CLI installed successfully.`);
