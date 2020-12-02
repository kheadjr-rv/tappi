const cacheName = "app-" + "0.0.12";

self.addEventListener("install", event => {
  console.log("installing app worker 0.0.12");
  self.skipWaiting();

  event.waitUntil(
    caches.open(cacheName).then(cache => {
      return cache.addAll([
        "/tappi",
        "/tappi/app.css",
        "/tappi/app.js",
        "/tappi/manifest.webmanifest",
        "/tappi/wasm_exec.js",
        "/tappi/web/app.wasm",
        "/tappi/web/css/docs.css",
        "/tappi/web/css/prism.css",
        "/tappi/web/images/maskable_icon_192px.png",
        "/tappi/web/images/maskable_icon_512px.png",
        "/tappi/web/js/prism.js",
        "https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500&display=swap",
        "https://fonts.googleapis.com/css2?family=Roboto&display=swap",
        
      ]);
    })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker 0.0.12 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
