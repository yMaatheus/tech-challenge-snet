<script setup lang="ts">
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
  DialogDescription,
  DialogTrigger,
  DialogClose,
} from "~/components/ui/dialog";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";
import { ref, reactive, watch } from "vue";
import type { Store } from "~/types/store";

const props = defineProps<{
  open?: boolean;
  title?: string;
  initialValues?: Partial<Store>;
  onSave: (form: Store) => void;
}>();

const emit = defineEmits(["update:open"]);
const localOpen = ref(props.open ?? false);

const form = reactive<Store>({
  id: '',
  number: "",
  name: "",
  corporate_name: "",
  address: "",
  address_number: "",
  city: "",
  state: "",
  zip_code: "",
});

watch(
  () => props.initialValues,
  (newValues) => {
    Object.assign(
      form,
      {
        id: '',
        number: "",
        name: "",
        corporate_name: "",
        address: "",
        address_number: "",
        city: "",
        state: "",
        zip_code: "",
      },
      newValues || {}
    );
  },
  { immediate: true }
);

watch(
  () => props.open,
  (val) => {
    localOpen.value = val ?? false;
  }
);

watch(localOpen, (val) => {
  emit("update:open", val);
});

function resetForm() {
  Object.assign(form, {
    id: '',
    number: "",
    name: "",
    corporate_name: "",
    address: "",
    address_number: "",
    city: "",
    state: "",
    zip_code: "",
  });
}


function submitForm() {
  props.onSave({ ...form });

  if (!props.initialValues?.id) {
    resetForm();
  }
}

</script>

<template>
  <Dialog v-model:open="localOpen">
    <DialogTrigger as-child>
      <slot name="trigger">
        <Button>Abrir formulário</Button>
      </slot>
    </DialogTrigger>
    <DialogContent>
      <DialogHeader>
        <DialogTitle>
          {{ props.title ?? "Cadastrar Loja" }}
        </DialogTitle>
        <DialogDescription class="sr-only">
          Preencha os campos obrigatórios.
        </DialogDescription>
      </DialogHeader>

      <form class="space-y-4 mt-4" @submit.prevent="submitForm">
        <div>
          <label class="block text-sm font-medium mb-1">Número da Loja</label>
          <Input
            v-model="form.number"
            type="text"
            class="input w-full"
            required
          />
        </div>
        <div>
          <label class="block text-sm font-medium mb-1">Nome</label>
          <Input
            v-model="form.name"
            type="text"
            class="input w-full"
            required
          />
        </div>
        <div>
          <label class="block text-sm font-medium mb-1">Razão Social</label>
          <Input
            v-model="form.corporate_name"
            type="text"
            class="input w-full"
            required
          />
        </div>
        <div>
          <label class="block text-sm font-medium mb-1">Endereço</label>
          <Input
            v-model="form.address"
            type="text"
            class="input w-full"
            required
          />
        </div>
        <div>
          <label class="block text-sm font-medium mb-1"
            >Número</label
          >
          <Input
            v-model="form.address_number"
            type="text"
            class="input w-full"
            required
          />
        </div>
        <div>
          <label class="block text-sm font-medium mb-1">Cidade</label>
          <Input
            v-model="form.city"
            type="text"
            class="input w-full"
            required
          />
        </div>
        <div>
          <label class="block text-sm font-medium mb-1">Estado</label>
          <Input
            v-model="form.state"
            type="text"
            maxlength="2"
            class="input w-full"
            required
          />
        </div>
        <div>
          <label class="block text-sm font-medium mb-1">CEP</label>
          <Input
            v-model="form.zip_code"
            type="text"
            class="input w-full"
            required
          />
        </div>
        <DialogFooter class="mt-6 flex justify-end gap-2">
          <DialogClose as-child>
            <Button
              type="button"
              variant="ghost"
              @click="emit('update:open', false)"
              >Cancelar</Button
            >
          </DialogClose>
          <Button type="submit">Salvar</Button>
        </DialogFooter>
      </form>
    </DialogContent>
  </Dialog>
</template>
