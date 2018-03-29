cmake_minimum_required(VERSION 3.2 FATAL_ERROR)

set(PROJECT eng_breaker_backend)
project(PROJECT CXX)

include(GNUInstallDirs)
enable_testing()

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)
add_subdirectory(app)

