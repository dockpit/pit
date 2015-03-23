var gulp = require('gulp'),
    browserify = require('browserify'),
    reactify = require('reactify'),
    source = require('vinyl-source-stream'),
    util = require('gulp-util'),
    watchify = require('watchify'),
    pkg = require('./package.json');
    
//build scripts js/jsx
gulp.task('jsx', function () {
  var bundler = browserify({
      entries: [pkg.paths.app], 
      transform: [reactify], 
      debug: true, 
      cache: {}, packageCache: {}, fullPaths: true
  });
  var watcher  = watchify(bundler);

  return watcher
    .on('update', function () {
        watcher.bundle() 
        .pipe(source(pkg.dest.app))        
        .pipe(gulp.dest(pkg.dest.dist));
        util.log("Rebuilt scripts");
    })
    .bundle() // Create the initial bundle when starting the task
    .pipe(source(pkg.dest.app))
    .pipe(gulp.dest(pkg.dest.dist));
});

gulp.task('default', ['jsx']);