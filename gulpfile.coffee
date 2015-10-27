gulp = require('gulp')
browserify = require('browserify')
tsify = require('tsify')
source = require 'vinyl-source-stream'

gulp.task 'build', ->
  browserify
    entries: ['./web/scripts/hello.ts']
  .plugin tsify
  .bundle()
  .pipe source 'bundle.js'
  .pipe gulp.dest './web/build'

