import htmx from "htmx.org";
import * as Flowbite from "flowbite";

import Alpine from "alpinejs";

// Declaration to avoid linting issues
declare global {
  interface Window {
    Alpine: typeof Alpine;
    htmx: typeof htmx;
  }
}

// Add global objects.
window.htmx = htmx;
window.Alpine = Alpine;

Alpine.start();

// Document ready function to ensure the DOM is fully loaded.
document.addEventListener("DOMContentLoaded", async () => {
  Flowbite.initFlowbite();

  // Add event listeners for HTMX events
  document.addEventListener("htmx:afterSwap", function () {
    Flowbite.initFlowbite();
  });
  document.addEventListener("htmx:afterRequest", function () {
    Flowbite.initFlowbite();
  });
  document.addEventListener("htmx:afterSettle", function () {
    Flowbite.initFlowbite();
  });

  await import("htmx-ext-response-targets");
});
