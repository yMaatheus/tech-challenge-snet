import type { Store } from "~/types/store";

type CreateStorePayload = {
  establishment_id: number;

  number: string;
  name: string;
  corporate_name: string;
  address: string;
  address_number: string;
  city: string;
  state: string;
  zip_code: string;
};

export async function createStore(payload: CreateStorePayload) {
  const config = useRuntimeConfig();
  const response = await $fetch<Store>(
    `${config.public.apiBase}/stores`,
    {
      method: "POST",
      body: {
        establishment_id: payload.establishment_id,

        number: payload.number,
        name: payload.name,
        corporate_name: payload.corporate_name,
        address: payload.address,
        address_number: payload.address_number,
        city: payload.city,
        state: payload.state,
        zip_code: payload.zip_code,
      },
    }
  );

  return response;
}
