add_executable(app
    src/main.cpp)
add_dependencies(app cpp_redis)
target_link_libraries (app cpp_redis)
