CXX			= g++
CXXFLAGS	= -Wall -Wextra -Werror -O3
SRCS		= $(wildcard *.cpp)
OBJS		= $(SRCS:.cpp=.o)
LDLIBS		= -lcurl -lpthread
PROGS		= $(basename $(SRCS))
RM			= rm -vf

.PHONY: all clean

all: $(PROGS)

clean:
	$(RM) $(OBJS) $(PROGS)
