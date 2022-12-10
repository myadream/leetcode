cmake_minimum_required(VERSION 3.17)
project(c C)

set(CMAKE_C_STANDARD 99)

add_subdirectory(src)

link_directories (${PROJECT_SOURCE_DIR}/src/tools)

