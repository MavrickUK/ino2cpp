# ino2cpp
[![Build Status](https://github.com/tj/commander.js/workflows/build/badge.svg)](https://github.com/tj/commander.js/actions?query=workflow%3A%22build%22)

Convert Arduino INO sketches to C++ by creating new .cpp and .h files


## Details

Arduino sketches and C++ are very similar.
However, an INO file cannot be compiled as-is by C/C++ compilers (e.g. GCC).
This tool converts INO sketches to C++ code such that off-the-shelf compilers and static analysis tools can be executed on the code.

## Converter
There are three steps in this conversion [1][2]:****
1. **Generate forward declarations**. Arduino INO sketches allow the use of a function before its definition. The first step is to parse the INO sketch to obtain function signatures, and generate a header file with these signatures ("sketch_name.h").
2. **Includes**. Two includes are inserted before the content of the INO sketch: #include <Arduino.h>, and #include "sketch_name.h".
3. **Write C++ to disk**. Write the resulting C++ and header file to disk. INO is unchanged!

## Installation

Portable, just download the latest release version and run the .exe from where you unzipped.

## Usage

```
ino2cpp -i [filename of .ino]
ino2cpp -h
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## References
[1] [How to convert Arduino example sketch to a complete c++ project?](https://arduino.stackexchange.com/questions/32998/how-to-convert-arduino-example-sketch-to-a-complete-c-project)
[2] [converting .ino files to .cpp - Arduino Forum](https://forum.arduino.cc/t/converting-ino-files-to-cpp/226366)

### License
[MIT](https://choosealicense.com/licenses/mit/)
****