<template>
  <div class="flex min-h-screen items-center justify-center bg-background">
    <form
      class="bg-card p-8 rounded-lg shadow-lg w-full max-w-sm space-y-6"
      @submit.prevent="login"
    >
      <h1 class="text-2xl font-bold text-center text-gray-700">
        Sistema de Gestão
      </h1>

      <Input
        v-model="email"
        type="email"
        placeholder="Seu e-mail"
        required
        class="w-full px-4 py-2"
      />

      <Input
        v-model="password"
        type="password"
        placeholder="Sua senha"
        required
        class="w-full px-4 py-2"
      />

      <Button
        type="submit"
        variant="default"
        size="lg"
        class="w-full font-semibold py-2 rounded-lg transition"
      >
        Login
      </Button>

      <FeedbackMessage
        v-if="errorMessage"
        :message="errorMessage"
        type="error"
      />
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";

definePageMeta({
  layout: false,
});

const email = ref("");
const password = ref("");
const errorMessage = ref("");
const router = useRouter();

function login() {
  if (!(email.value && password.value)) {
    errorMessage.value = "Please enter email and password.";
    return;
  }

  localStorage.setItem("logged", "true");
  router.push("/dashboard");
}
</script>
