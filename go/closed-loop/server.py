import signal
from socket import socket, AF_INET, SOCK_DGRAM
from threading import Thread

def signal_handler(signal, frame):
    global interrupted
    interrupted = True

signal.signal(signal.SIGINT, signal_handler)

interrupted = False

def safe_print(content):
    '''
    >>> import dis
    >>> def foo(x): print x
    >>> dis.dis(foo) # PRINT_NEWLINE is a separate instruction
    '''
    print "{0}\n".format(content),

def server(address):
    sock = socket(AF_INET, SOCK_DGRAM)
    sock.bind(address)
    while True:
        msg, addr = sock.recvfrom(8192)
        safe_print(str(addr) + " : " + str(msg))
        if interrupted:
            break

if __name__ == '__main__':
    host = ''
    ports = [2561, 2563, 2564, 2565, 2569]
    threads = []
    for port in ports:
        address = (host, port)
        t = Thread(target=server, args=(address,))
        threads.append(t)
        t.start()
