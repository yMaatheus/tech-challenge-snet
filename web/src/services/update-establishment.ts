import type { Establishment } from "~/types/establishment"

export async function updateEstablishment(id: string | number, payload: Partial<Establishment>) {
  const config = useRuntimeConfig()
  const response = await $fetch<Establishment>(
    `${config.public.apiBase}/establishments/${id}`,
    {
      method: 'PUT',
      body: {
        number: payload.number,
        name: payload.name,
        corporate_name: payload.corporate_name,
        address: payload.address,
        address_number: payload.address_number,
        city: payload.city,
        state: payload.state,
        zip_code: payload.zip_code,
      }
    }
  )

  return response
}