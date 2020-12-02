const cacheName = "app-" + "0.0.15";

self.addEventListener("install", event => {
  console.log("installing app worker 0.0.15");
  self.skipWaiting();

  event.waitUntil(
    caches.open(cacheName).then(cache => {
      return cache.addAll([
        "/",
        "/app.css",
        "/app.js",
        "/manifest.webmanifest",
        "/wasm_exec.js",
        "/web/app.wasm",
        "/web/css/docs.css",
        "/web/css/prism.css",
        "/web/documents/architecture.md",
        "/web/documents/start.md",
        "/web/documents/terraform.md",
        "/web/images/maskable_icon_192px.png",
        "/web/images/maskable_icon_512px.png",
        "/web/js/prism.js",
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
  console.log("app worker 0.0.15 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
