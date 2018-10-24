import json
import logging
import logging.config
import signal
from socket import socket, AF_INET, SOCK_DGRAM
from threading import Thread

class StructuredMessage(object):
    def __init__(self, **kwargs):
        self.kwargs = kwargs

    def __str__(self):
        return '%s' % (json.dumps(self.kwargs),)

_ = StructuredMessage   # optional, to improve readability

# logging.basicConfig(level=logging.INFO, format='%(message)s')
# logging.info(_('message 1', foo='bar', bar='baz', num=123, fnum=123.456))

def server(address):
    sock = socket(AF_INET, SOCK_DGRAM)
    sock.bind(address)
    logger = logging.getLogger()
    while True:
        msg, addr = sock.recvfrom(8192)
        # safe_print(str(addr) + " : " + str(msg))
        logger.debug(_(addr=addr, msg=msg))

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
