<template>
  <div class="container mx-auto py-8 px-4">
    <div
      class="flex flex-col items-center gap-5 justify-between mb-6 md:flex-row"
    >
      <h1 class="text-2xl font-bold">Estabelecimentos</h1>

      <EstablishmentForm
        v-model:open="isDialogOpen"
        title="Cadastrar Estabelecimento"
        :on-save="handleCreateEstablishment"
      >
        <template #trigger>
          <Button variant="default" size="lg" class="w-3/4 md:w-auto" @click="isDialogOpen = true">
            <Plus class="w-4 h-4" />
            Criar Estabelecimento
          </Button>
        </template>
      </EstablishmentForm>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="e in establishments"
        :key="e.id"
        class="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow"
      >
        <h3 class="text-lg font-semibold text-gray-900 mb-2">
          {{ e.name }}
        </h3>

        <p class="text-gray-600 mb-4">
          {{
            e.address + " - " + e.city + " - " + e.state + " - " + e.zip_code
          }}
        </p>
        <p class="text-sm text-gray-500 mb-4">
          {{ e.storesTotal }} loja(s) vinculada(s)
        </p>

        <div class="flex space-x-2">
          <Button class="flex-1" as-child>
            <NuxtLink :to="`/dashboard/${e.id}`"> Ver Lojas </NuxtLink>
          </Button>

          <EstablishmentForm
            title="Editar Estabelecimento"
            :on-save="handleUpdateEstablishment"
            :initial-values="e"
          >
            <template #trigger>
              <Button class="flex-1 bg-yellow-600 hover:bg-yellow-700">
                Editar
              </Button>
            </template>
          </EstablishmentForm>

          <Button
            class="flex-1 bg-red-600 hover:bg-red-700"
            @click="handleDeleteEstablishment(e)"
          >
            Excluir
          </Button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { Button } from "~/components/ui/button";
import type { Establishment } from "~/types/establishment";
import { Plus } from "lucide-vue-next";
import { deleteEstablishment } from "~/services/delete-establishment";
import { fetchEstablishments } from "~/services/fetch-establishments";
import EstablishmentForm from "~/components/app/EstablishmentForm.vue";
import { createEstablishment } from "~/services/create-establishment";
import { updateEstablishment } from "~/services/update-establishment";
import { toast } from "vue-sonner";

definePageMeta({
  middleware: "auth",
});

const isDialogOpen = ref(false);

const { data: establishments, refresh } = await fetchEstablishments();

async function handleCreateEstablishment(data: Establishment) {
  try {
    await createEstablishment(data);

    isDialogOpen.value = false;
    refresh();

    toast.success("Estabelecimento foi criado com sucesso!", {
      description: `${data.name} - ${data.address}, ${data.address_number}`,
    });
  } catch (error) {
    if (error instanceof Error) {
      toast.error("Erro ao tentar criar o estabelecimento.");
    }
  }
}

async function handleUpdateEstablishment(data: Establishment) {
  try {
    if (!data.id) {
      throw new Error("ID do estabelecimento é necessário para atualização.");
    }

    await updateEstablishment(data.id, data);
    refresh();

    toast.success("Estabelecimento foi atualizado com sucesso!", {
      description: `${data.name} - ${data.address}, ${data.address_number}`,
    });
  } catch (error) {
    if (error instanceof Error) {
      toast.error("Erro ao tentar atualizar o estabelecimento.");
    }
  }
}

async function handleDeleteEstablishment(establishment: Establishment) {
  if (confirm("Você quer remover este estabelecimento?")) {
    try {
      await deleteEstablishment(establishment.id);
      await refresh();

      toast.success("Estabelecimento removido com sucesso!", {
        description: `${establishment.name}`,
      });
    } catch (error) {
      if (error instanceof Error) {
        toast.error("Erro ao tentar remover o estabelecimento.");
      }
    }
  }
}
</script>
