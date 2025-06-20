export default defineNuxtRouteMiddleware(() => {
  if (!import.meta.client) return;

  if (!localStorage.getItem("logged")) {
    return navigateTo("/");
  }
});
