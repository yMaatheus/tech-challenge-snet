import type { EstablishmentWithStores } from "~/types/establishment";

export async function fetchEstablishmentById(establishmentId: string) {
  const config = useRuntimeConfig();

  const response = await useFetch<EstablishmentWithStores>(
    `${config.public.apiBase}/establishments/${establishmentId}`,
    {
      method: "GET",
      headers: { "Content-Type": "application/json" },
    }
  );
  return response;
}
