include(ExternalProject)

cmake_minimum_required(VERSION 3.2 FATAL_ERROR)
project(Backend VERSION 0.1 LANGUAGES CXX)
include(GNUInstallDirs)
enable_testing()
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)
add_subdirectory(app)
configure_file(CppRedis.CMakeLists.txt cpp-redis-download/CMakeLists.txt)
execute_process(COMMAND "${CMAKE_COMMAND}" -G "${CMAKE_GENERATOR}" .
    WORKING_DIRECTORY "${CMAKE_BINARY_DIR}/cpp-redis-download")
execute_process(COMMAND "${CMAKE_COMMAND}" --build .
    WORKING_DIRECTORY "${CMAKE_BINARY_DIR}/cpp-redis-download" )
add_subdirectory("${CMAKE_BINARY_DIR}/cpp-redis-src"
                 "${CMAKE_BINARY_DIR}/cpp-redis-build")


