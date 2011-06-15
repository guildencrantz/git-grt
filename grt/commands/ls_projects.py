from grt.GerritClient import GerritClient

__author__        = 'Matt Henkel'
__copyright__    = 'Copyright 2011, Return Path, Inc.'
__credits__        = ['Matt Henkel']
__license__        = 'Apache License'
__version__        = '2.0'
__maintainer__    = 'Matt Henkel'
__email__        = 'mhenkel@returnpath.net'

class ls_projects(object):
    Description = 'Get a list of gerrit projects you have read access to.'

    def __init__(self, config):
        self.config = config

    def Do(self, commandArgs = None):
        command = 'ls-projects'

        gerritClient = GerritClient(self.config)
        gerritClient.connect()
        (stdout, stderror) = gerritClient.execute(command)
        # TODO: Sufficient error checking?
        if len(stdout) > 0:
            return stdout
        else:
            raise Exception("There was an error executing %s: %s" % (command, stderror))

# vim: ts=4 sw=4 sts=4 et ai ft=python :
