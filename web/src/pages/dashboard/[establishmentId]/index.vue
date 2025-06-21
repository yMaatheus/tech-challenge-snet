<template>
  <div class="container mx-auto py-8 px-4">
    <div
      class="flex gap-5 flex-col items-center justify-between mb-6 md:flex-row"
    >
      <div class="flex flex-col space-x-4 space-y-2 w-full md:flex-row">
        <NuxtLink :to="`/dashboard`" as-child>
          <Button
            variant="ghost"
            class="text-blue-600 hover:text-blue-800 text-center"
          >
            <ChevronLeft class="w-4 h-4" />
            Voltar
          </Button>
        </NuxtLink>

        <h2 class="text-center text-2xl font-bold text-gray-900 md:text-left">
          Lojas - {{ establishment?.name || "Carregando..." }}
        </h2>
      </div>

      <StoreForm
        v-model:open="isDialogOpen"
        title="Criar Loja"
        :on-save="handleCreateStore"
      >
        <template #trigger>
          <Button
            variant="default"
            size="lg"
            class="w-3/4 md:w-auto"
            @click="isDialogOpen = true"
          >
            <Plus class="w-4 h-4" />
            Nova Loja
          </Button>
        </template>
      </StoreForm>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="store in establishment?.stores || []"
        :key="store.id"
        class="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow"
      >
        <h3 class="text-lg font-semibold text-gray-900 mb-2">
          {{ store.name }}
        </h3>
        <p class="text-gray-600 mb-2">
          {{
            store.address +
            " - " +
            store.city +
            " - " +
            store.state +
            " - " +
            store.zip_code
          }}
        </p>
        <p class="text-sm text-gray-500 mb-4">Número: {{ store.number }}</p>

        <div class="flex space-x-2">
          <StoreForm
            title="Editar Loja"
            :on-save="handleUpdateEstablishment"
            :initial-values="store"
          >
            <template #trigger>
              <Button
                variant="default"
                class="flex-1 bg-yellow-600 hover:bg-yellow-700"
              >
                Editar
              </Button>
            </template>
          </StoreForm>

          <Button
            variant="destructive"
            class="flex-1 bg-red-600 hover:bg-red-700"
            @click="handleDeleteStore(store)"
          >
            Excluir
          </Button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ChevronLeft, Plus } from "lucide-vue-next";
import { toast } from "vue-sonner";
import StoreForm from "~/components/app/StoreForm.vue";
import { Button } from "~/components/ui/button";
import { createStore } from "~/services/create-store";
import { deleteStore } from "~/services/delete-store";
import { fetchEstablishmentById } from "~/services/fetch-establishment-by-id";
import { updateStore } from "~/services/update-store";
import type { Store } from "~/types/store";

const route = useRoute();
const establishmentId = route.params.establishmentId as string;

const isDialogOpen = ref(false);

const { data: establishment, refresh } = await fetchEstablishmentById(
  establishmentId
);

async function handleCreateStore(data: Store) {
  try {
    await createStore({
      ...data,
      establishment_id: Number(establishmentId),
    });
    await refresh();

    isDialogOpen.value = false;

    toast.success("Loja foi criada com sucesso!", {
      description: `${data.name} - ${data.address}, ${data.address_number}`,
    });
  } catch (error) {
    if (error instanceof Error) {
      toast.error("Erro ao tentar criar a loja.");
    }
  }
}

async function handleUpdateEstablishment(data: Store) {
  try {
    await updateStore(data.id, {
      ...data,
      establishment_id: Number(establishmentId),
    });
    await refresh();

    toast.success("Loja foi atualizada com sucesso!", {
      description: `${data.name} - ${data.address}, ${data.address_number}`,
    });
  } catch (error) {
    if (error instanceof Error) {
      toast.error("Erro ao tentar atualizar a loja.");
    }
  }
}

async function handleDeleteStore(store: Store) {
  if (confirm("Você quer mesmo remover essa loja?")) {
    try {
      await deleteStore(store.id);
      await refresh();

      toast.success("Loja removida com sucesso!", {
        description: `${store.name}`,
      });
    } catch (error) {
      if (error instanceof Error) {
        toast.error("Erro ao tentar remover a loja.");
      }
    }
  }
}
</script>
