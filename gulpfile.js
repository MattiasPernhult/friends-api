var gulp = require('gulp');
var path = require('path');
var shell = require('gulp-shell');

var goPath = 'source/**/*.go';

gulp.task('compilepkg', function() {
  return gulp.src(goPath, {
      read: false
    })
    .pipe(shell(['go install <%= stripPath(file.path) %>'], {
      templateData: {
        stripPath: function(filePath) {
          var subPath = filePath.substring(process.cwd().length);
          var pkg = subPath.substring(1, subPath.lastIndexOf(path.sep));
          return 'sandbox/friends-api/' + pkg;
        }
      }
    }));
});

gulp.task('watch', function() {
  gulp.watch(goPath, ['compilepkg']);
});
