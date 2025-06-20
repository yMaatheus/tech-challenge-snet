import type { EstablishmentWithStoresTotal } from "~/types/establishment";

export async function fetchEstablishments() {
  const config = useRuntimeConfig();

  const response = await useFetch<EstablishmentWithStoresTotal[]>(
    `${config.public.apiBase}/establishments`,
    {
      method: "GET",
      headers: { "Content-Type": "application/json" },
    }
  );
  return response;
}
