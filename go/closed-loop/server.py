import logging
import logging.config
import signal
from socket import socket, AF_INET, SOCK_DGRAM
from threading import Thread

def server(address):
    sock = socket(AF_INET, SOCK_DGRAM)
    sock.bind(address)
    logger = logging.getLogger()
    while True:
        msg, addr = sock.recvfrom(8192)
        # safe_print(str(addr) + " : " + str(msg))
        logger.debug("%s : %s", addr, msg)

if __name__ == '__main__':
    logging.config.fileConfig('logging_config.ini')
    host = ''
    ports = [2561, 2563, 2564, 2565, 2569]
    threads = []
    for port in ports:
        address = (host, port)
        t = Thread(target=server, args=(address,))
        threads.append(t)
        t.start()
