import scp
import paramiko
import getpass

__author__     = 'Matt Henkel'
__copyright__  = 'Copyright 2011, Return Path, Inc.'
__credits__    = ['Matt Henkel']
__license__    = 'Apache License'
__version__    = '2.0'
__maintainer__ = 'Matt Henkel'
__email__      = 'mhenkel@returnpath.net'

class GerritClient:
    def __init__(self, config):
        self.gerrit_host = config.get_value('gerrit', 'host')
        # get_value returns a long, but we require an integer.
        self.gerrit_ssh_port = int(config.get_value('gerrit', 'ssh-port', 29418))
        self.gerrit_username = config.get_value('gerrit', 'username', getpass.getuser())
        self.client = paramiko.SSHClient()
        # TODO: This isn't the safest thing: It just connects to any host key
        self.client.set_missing_host_key_policy(paramiko.AutoAddPolicy())

    def connect(self):
        self.client.connect(self.gerrit_host, port=self.gerrit_ssh_port, username=self.gerrit_username, timeout=10)

    def close(self):
        self.client.close()

    def execute(self, command):
        (stdin, stdout, stderr) = self.client.exec_command('gerrit ' + command)
        stdout = stdout.readlines()
        stderr = stderr.readlines()
        self.close()
        return stdout, stderr

    def fetch(self, remotePath, localPath):
        self.connect()
        scpClient = scp.SCPClient(self.client._transport)
        scpClient.get(remotePath, localPath)
        self.close()

# vim: set ts=4 sw=4 sts=4 et ai :
