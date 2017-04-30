# Add your own tasks in files placed in lib/tasks ending in .rake,
# for example lib/tasks/capistrano.rake, and they will automatically be available to Rake.

require_relative 'config/application'

Rails.application.load_tasks

@npm_dir = ENV["NPM_DIR"] || "templates/open-digit"
namespace :npm do
  task :build do
    sh "cd #{@npm_dir} && npm run build" do |res, ok|
      "Failed to run npm build" unless ok
    end
  end

  task :serve do
    sh "cd #{@npm_dir} && npm run serve:dist" do |res, ok|
      "Failed to run npm serve" unless ok
    end
  end
end

namespace :gulp do
  task :build do
    sh "cd #{@npm_dir} && gulp build" do |res, ok|
      "Failed to run gulp build" unless ok
    end
  end

  task :serve do
    sh "cd #{@npm_dir} && gulp serve" do |res, ok|
      "Failed to run gulp serve" unless ok
    end
  end
end

task :run do
  sh "rails s -p3030"
end
