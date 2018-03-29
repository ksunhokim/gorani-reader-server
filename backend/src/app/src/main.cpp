#include <iostream>
#include <cpp_redis/cpp_redis>
#include <SimpleAmqpClient/SimpleAmqpClient.h>

int main(int argc, char *argv[]) {
	AmqpClient::Channel::ptr_t connection = AmqpClient::Channel::Create("localhost");	
}
