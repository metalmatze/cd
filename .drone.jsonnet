local pipeline = {
  kind: 'pipeline',
  name: 'build',
  platform: {
    os: 'linux',
    arch: 'amd64',
  },
  steps: [],
};

local golang = {
  name: 'golang',
  image: 'golang:1.12',
  pull: 'always',
  environment: {
    CGO_ENABLED: '0',
    GO111MODULE: 'on',
    GOPROXY: 'https://proxy.golang.org',
  },
};

local docker = {
  name: 'docker',
  image: 'plugins/docker',
  settings+: {
    registry: 'quay.io',
    repo: 'quay.io/%s',
    username: {
      from_secret: 'quay_username',
    },
    password: {
      from_secret: 'quay_password',
    },
  },
  when: {
    branch: [
      'master',
    ],
  },
};

[
  pipeline {
    steps+: [
      golang {
        name: 'build',
        commands: [
          'make build',
        ],
      },
      {
        name: 'dart',
        image: 'google/dart:2.3',
        pull: 'always',
        commands: [
          'cd ui',
          'pub get --no-precompile',
          'pub global activate webdev',
          '~/.pub-cache/bin/webdev build',
          'rm -rf build/packages',
        ],
      },
    ] + [
      docker {
        name: 'docker-%s' % name,
        settings+: {
          repo: 'quay.io/signalcd/%s' % name,
          dockerfile: 'cmd/%s/Dockerfile' % name,
          context: './cmd/%s/' % name,
        },
      }
      for name in ['api', 'agent']
    ],
  },

  pipeline {
    name: 'examples',
    steps+: [
      docker {
        name: 'docker-examples-%s' % name,
        settings+: {
          repo: 'quay.io/signalcd/examples:%s' % name,
          dockerfile: 'examples/%s/Dockerfile' % name,
          context: 'examples/%s' % name,
        },
      }
      for name in ['cheese0', 'cheese1', 'cheese2']
    ],
  },
]
