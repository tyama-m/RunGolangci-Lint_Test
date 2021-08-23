//const fs = require('fs');
//const path = 'main.go';
//if (fs.existsSync(path)) {
//    console.log('ファイル・ディレクトリは存在します。');
//} else {
//    console.log('ファイル・ディレクトリは存在しません。');
//}

var fs = require('fs');
fs.readdir('.', function(err, files){
    if (err) throw err;
    var fileList = files.filter(function(file){
        return fs.statSync(file).isFile() && /.*\.*$/.test(file); //絞り込み
    })
    console.log(fileList);
});