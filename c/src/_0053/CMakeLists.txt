
link_libraries(tools)
add_executable(_0053 main.c maxSubArray.c maxSubArray.h)
target_link_libraries(_0053 tools)
install(TARGETS tools LIBRARY DESTINATION lib)