cmake_minimum_required(VERSION 3.15)
project(libluau CXX)

add_library(luau SHARED luau_bridge.cpp)
target_compile_features(luau PRIVATE cxx_std_17)