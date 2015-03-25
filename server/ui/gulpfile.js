var gulp = require('gulp'),
    browserify = require('browserify'),
    reactify = require('reactify'),
    source = require('vinyl-source-stream'),
    util = require('gulp-util'),
    watchify = require('watchify'),
    pkg = require('./package.json'),
    less = require('gulp-less'),
    plumber = require('gulp-plumber');
    

gulp.task("less", function () {
  gulp.src(pkg.paths.less)
    .pipe(plumber())
    .pipe(less())
    .pipe(gulp.dest(pkg.dest.dist));
});

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
    .bundle()
    .pipe(source(pkg.dest.app))
    .pipe(gulp.dest(pkg.dest.dist));
});

gulp.task('default', ['jsx'], function () {
  gulp.watch(pkg.paths.less, ['less']);
});