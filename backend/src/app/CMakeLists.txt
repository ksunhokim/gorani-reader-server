add_executable(app
    src/main.cpp)
target_link_libraries(app cpp_redis SimpleAmqpClient)
