<template>
  <div class="min-h-screen bg-background">
    <header class="bg-card shadow-sm">
      <div class="mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-16">
          <div class="flex items-center space-x-4">
            <h1 class="text-xl font-semibold text-gray-900">
              Sistema de Gestão
            </h1>
          </div>
          <Button variant="link" @click="logout"> Sair </Button>
        </div>
      </div>
    </header>

    <NuxtPage />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { Button } from "@/components/ui/button";

const isLogged = ref(false);
const router = useRouter();

onMounted(() => {
  isLogged.value = !!localStorage.getItem("logged");
});

function logout() {
  localStorage.removeItem("logged");
  router.push("/");
}
</script>
