include_directories(${CMAKE_LIBRARY_PATH})

add_executable(_0053 main.c maxSubArray.c maxSubArray.h)

target_link_libraries(_0053 Tools)