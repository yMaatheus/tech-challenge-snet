import type { Establishment } from "~/types/establishment";

export async function fetchEstablishments() {
  const config = useRuntimeConfig();

  const response = await useFetch<Establishment[]>(
    `${config.public.apiBase}/establishments`,
    {
      method: "GET",
      headers: { "Content-Type": "application/json" },
    }
  );
  return response;
}
