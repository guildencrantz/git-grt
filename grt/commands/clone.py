from grt.GerritClient import GerritClient
from grt.commands.ls_projects import ls_projects
# GitPython
from git import *

__author__		= 'Matt Henkel'
__copyright__	= 'Copyright 2011, Return Path, Inc.'
__credits__		= ['Matt Henkel']
__license__		= 'Apache License'
__version__		= '2.0'
__maintainer__	= 'Matt Henkel'
__email__		= 'mhenkel@returnpath.net'

class clone(object):
	Description = 'Clone a repository from gerrit.'

	def __init__(self, config):
		self.config = config

	def Do(self, projects=None):
		if not projects or not len(projects):
			projects = self.SelectProject()
			if projects.endswith('\n'):
				projects = projects[:-1]

		# We're going to assume the user intends to clone multiple repositories
		for project in projects:
			self.CloneProject(project)

	def SelectProject(self):
		projects = ls_projects(self.config).Do()
		# We need to make sure we pad the correct number of leading zeroes no matter how many projects exist
		list_format = '%0'+str(len(str(len(projects))))+'d.     %s'
		print ''.join(map(lambda x: list_format % (projects.index(x)+1, x), projects))
		projectNumber = raw_input('Enter the number for the project you\'d like to clone: ')
		if projectNumber.isdigit():
			projectNumber = int(projectNumber)
			if projectNumber > 0 and projectNumber <= len(projects):
				return projects[int(projectNumber) - 1]
			else:
				print 'Valid project numbers are between (inclusive) 1 and %d' % (len(projects))
		else:
			print 'You must enter a number between 1 and %d' % (len(projects))
		#FIXME: recursiveize this

	def CloneProject(self, project):
		result = Git().clone("ssh://%s:%s/%s" % (self.config.get_value('gerrit', 'host'), self.config.get_value('gerrit', 'ssh-port', 29418), project))
		print result
		repo = Repo(project)
		print "Configuring gerrit remote"
		repo.remotes.origin.rename('gerrit')
		config_writer = repo.config_writer().set('branch "master"', 'push', 'refs/for/master')
		print "Retrieving commit-msg hook"
		GerritClient(self.config).fetch('hooks/commit-msg', '%s/.git/hooks/' % project)
