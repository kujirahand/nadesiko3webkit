// nadesiko3webkit
// index.htmlで使うJavaScriptの関数を定義

var nako3_info_id = 1

function qs(query) {  
  return document.querySelector(query)
}

function to_html(s) {
  s = '' + s
  return s.replace(/\&/g, '&amp;')
          .replace(/\</g, '&lt;')
          .replace(/\>/g, '&gt;')
}

var nako3_get_resultbox = function () {
  return qs("#nako3result_div_" + nako3_info_id)
}
var nako3_get_info = function () {
  return qs("#nako3_info_" + nako3_info_id)
}
var nako3_get_error = function () {
  return qs("#nako3_error_" + nako3_info_id)
}
var nako3_get_canvas = function () {
  return qs("#nako3_canvas_" + nako3_info_id)
}
var nako3_get_div = function () {
  return qs("#nako3_div_" + nako3_info_id)
}
// 表示
var nako3_print = function (s) {
  console.log("[表示] " + s)
  var info = nako3_get_info()
  if (!info) return
  var box = nako3_get_resultbox()
  box.style.display = 'block'
  s = "" + s // 文字列に変換
  // エラーだった場合
  if (s.substr(0, 9) == "==ERROR==") {
    s = s.substr(9)
    var err = nako3_get_error()
    err.innerHTML = s
    err.style.display = 'block'
    return
  } else {
    info.innerHTML += to_html(s).replace(/\n/, "\n<br>") + "<br>\n"
    info.style.display = 'block'
  }
}
//---------------------------------
var nako3_clear = function (s, use_canvas) {
  var info = nako3_get_info()
  if (!info) return
  info.innerHTML = ''
  info.style.display = 'none'
  var err = nako3_get_error()
  err.innerHTML = ''
  err.style.display = 'none'
  var div = nako3_get_div()
  if (div) div.innerHTML = ''
  if (use_canvas) {
    var canvas = nako3_get_canvas()
    if (canvas) {
      var ctx = canvas.getContext('2d')
      ctx.clearRect(0, 0, canvas.width, canvas.height)
    }
  }
  if (navigator.nako3) {
    navigator.nako3.clearPlugins()
  }
}

//---------------------------------
function makePostData(params, sys) {
  const flist = []
  for (const key in params) {
    const v = params[key]
    const kv = encodeURIComponent(key) + '=' + encodeURIComponent(v)
    flist.push(kv)
  }
  return flist.join('&')
}

function sendPost(url, params, callback, sys) {
  const bodyData = makePostData(params)
  const options = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    body: bodyData
  }
  fetch(url, options).then(res => {
    return res.json()
  })
  .then(data => callback(data))
  .catch(err => {
    console.log(err)
    sys.__v0['AJAX:ONERROR'](err)
  })
}

function nako3_fileSave(cb, value, name, sys) {
  sendPost('/api', {
    func: 'save',
    name: name,
    value: value
  }, data => {
    console.log(data)
    cb(data)
  }, sys)
}

function nako3_fileLoad(cb, name, sys) {
  sendPost('/api', {
    func: 'load',
    name: name,
  }, data => {
    console.log(data)
    sys.__v0['対象'] = data.value
    cb(data)
  }, sys)
}

function nako3_files(cb, sys) {
  sendPost('/api', {
    func: 'files',
  }, data => {
    console.log(data)
    sys.__v0['対象'] = data.value.split(',')
    cb(data)
  }, sys)
}

//---------------------------------
// 独自関数の登録
const nako3_add_func = function () {
  //
  // この関数の中で定義した関数
  //
  const nako3 = navigator.nako3
  nako3.addFunc("表示", [['の', 'を', 'と']], nako3_print, true, false)
  nako3.addFunc("コンソール表示", [['の', 'を', 'と']], (s) => console.log(s), true, false)
  nako3.addFunc("表示ログクリア", [], nako3_clear, true, false)
  //
  // webkit/lorcaにバインドされたAPI --- callback
  //
  nako3.addFunc("ファイル保存時", [['で'], ['を'], ['へ', 'に']], 
    (cb, value, name, sys) => Nako3api_save(name, value).then(r => { sys.__v0['対象'] = r; cb(r) }, true, false))
  nako3.addFunc("ファイル読時", [['で'], ['を', 'の', 'から']], 
    (cb, name, sys) => Nako3api_load(name).then(r => { sys.__v0['対象'] = r; cb(r) }), false)
  nako3.addFunc("ファイル一覧取得時", [['で']], 
    (cb, sys) => Nako3api_files().then(r => { sys.__v0['対象'] = r; cb(r) }), false)
  nako3.addFunc("起動時", [['の'], ['を', 'で']], 
    (cb, path, sys) => {
      const args = typeof(path) == 'string' ? [path] : path
      Nako3api_exec(args).then(r => {
        sys.__v0['対象'] = r;
        cb(r) 
      })
    }, true)
  nako3.addFunc("環境変数取得時", [['で'], ['の']], 
    (cb, key, sys) => Nako3api_getenv(key).then(r => { sys.__v0['対象'] = r; cb(r) }), true)
  nako3.addFunc("環境変数設定時", [['で'], ['に','へ'], ['を']], 
    (cb, key, val) => Nako3api_setenv(key, val).then(r => cb(r)), true)
  nako3.addFunc("環境変数一覧取得時", [['で']], 
    (cb, sys) => Nako3api_envlist().then(results => {
      const obj = {}
      for (let line of results) {
        const a = line.split('=', 2)
        obj[a[0]] = a[1]
      }
      sys.__v0['対象'] = obj;
      cb(obj) 
  }), true)
  nako3.addFunc("内部情報取得時", [['で']], 
    (cb, sys) => Nako3api_info().then(r => {
      console.log('@@@', r);
      sys.__v0['対象'] = JSON.parse(r); cb(r) }), false)
  //
  // 非同期関数版(asyncFn)
  //
  nako3.addFunc("ファイル保存", [['を'], ['へ', 'に']], async (value, name, _sys) => await Nako3api_save(name, value), true, true)
  nako3.addFunc("ファイル読", [['を', 'の', 'から']], async (name, _sys) => await Nako3api_load(name), false, true)
  nako3.addFunc("ファイル一覧取得", [], async (_sys) => await Nako3api_files(), false, true)
  nako3.addFunc("起動", [['を', 'で']], async (path, _sys) => await Nako3api_exec(path), false, true)
  nako3.addFunc("環境変数取得", [['の']], async (key, _sys) => await Nako3api_getenv(key), false, true)
  nako3.addFunc("環境変数設定", [['に','へ'], ['を']], async (key, val) => await Nako3api_setenv(key, val), true, true)
  nako3.addFunc("環境変数一覧取得", [],
    async (_sys) => {
      const rawEnvList = await Nako3api_envlist()
      const obj = {}
      for (let line of rawEnvList) {
        const a = line.split('=', 2)
        obj[a[0]] = a[1]
      }
      console.log(obj)
      return obj
    }, false, true)
  nako3.addFunc("内部情報取得", [], async (_sys) => await Nako3api_info(), false, true)
}
//---------------------------------

