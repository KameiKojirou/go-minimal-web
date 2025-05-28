<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="drawer drawer-end">
    <input id="my-drawer" type="checkbox" class="drawer-toggle" />
    <div class="drawer-content">
      <!-- Page content here -->
       <div class="flex flex-row justify-center">
         <nav class="flex flex-row flex-1 gap-2 max-w-5xl justify-between [&>a]:btn [&>a]:btn-ghost">
           <div class="flex flex-row flex-1 gap-2">
             <RouterLink to="/">Home</RouterLink>
             <RouterLink to="/about">About</RouterLink>
            </div>
            <label for="my-drawer" class="btn btn-primary drawer-button">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 16 16">
                <path
                fill="currentColor"
                d="M1 2.75A.75.75 0 0 1 1.75 2h12.5a.75.75 0 0 1 0 1.5H1.75A.75.75 0 0 1 1 2.75m0 5A.75.75 0 0 1 1.75 7h12.5a.75.75 0 0 1 0 1.5H1.75A.75.75 0 0 1 1 7.75M1.75 12h12.5a.75.75 0 0 1 0 1.5H1.75a.75.75 0 0 1 0-1.5"
                />
              </svg>
            </label>
          </nav>
        </div>
      <main class="p-4">
        <RouterView />
      </main>
  </div>
  <div class="drawer-side">
    <label for="my-drawer" aria-label="close sidebar" class="drawer-overlay"></label>
    <ul class="menu bg-base-200 text-base-content min-h-full w-80 p-4">
      <!-- Sidebar content here -->
      <li><a>Sidebar Item 1</a></li>
      <li><a>Sidebar Item 2</a></li>
    </ul>
  </div>
</div>
</template>
