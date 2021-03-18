self.addEventListener("activate", function (event) { const current = ["rundocs-2.0.9"]; event.waitUntil( caches.keys().then(function (keyList) { return Promise.all( keyList.map(function (key) { if (current.indexOf(key) === -1) { return caches.delete(key); } }) ); }) ); }); self.addEventListener("fetch", function (e) { if (e.request.url.match("rundocs/jekyll-rtd-theme@2.0.9")) { e.respondWith( caches.match(e.request).then(function (resp) { if (resp !== undefined) { return resp; } else { return fetch(e.request, { cache: "no-store", }) .then(function (resp) { let clone = resp.clone(); caches.open("rundocs-2.0.9").then(function (cache) { cache.put(e.request, clone); }); return resp; }) .catch(console.log); } }) ); } });