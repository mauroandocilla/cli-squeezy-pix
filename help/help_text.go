package help

const RootShortHelpText = "SqueezyPix: A CLI for optimizing, converting, and generating responsive images for the web"
const RootHelpText = `SqueezyPix is a powerful and flexible image optimization CLI tool.

Description:
  SqueezyPix is a command-line tool that helps you optimize, convert, and generate responsive images 
  for the web. It supports multiple formats and provides advanced options for image quality 
  and file size reduction.

Examples:
  # Display help for a specific command
  spix [command] --help
`

const OptimizeShortHelpText = "Optimize an image with resizing and quality adjustments"
const OptimizeHelpText = `Optimize an image by resizing it to a specified size and adjusting the quality level.

Examples:
  spix opt -i input.jpg -o output.jpg -q 85  # Using aliases
  spix optimize --input input.jpg --output output.jpg --quality 85

Options:
  --input (-i): Specifies the input image file.
  --output (-o): Specifies the output image file.
  --quality (-q): Adjust the image quality (default: 85).
`

const ConvertShortHelpText = "Convert an image to another format"
const ConvertHelpText = `Convert an image to another format such as jpeg, png, or webp.

Examples:
  spix conv -i input.jpg -o output.png -f png  # Using aliases
  spix convert --input input.jpg --output output.png --format png

Options:
  --input (-i): Specifies the input image file.
  --output (-o): Specifies the output image file.
`

const ResponsiveShortHelpText = "Generate responsive images in multiple sizes"
const ResponsiveHelpText = `Generate responsive images for web development, with multiple sizes to support different breakpoints.

Examples:
  spix resp -i input.jpg -o output -w 320,640,1024  # Using aliases
  spix responsive --input input.jpg --output output --widths 320,640,1024

Options:
  --input (-i): Specifies the input image file.
  --output (-o): Specifies the base path for output images.
  --widths (-w): A list of widths for the responsive images.
`
