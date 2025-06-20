export async function deleteEstablishment(id: string): Promise<void> {
  const config = useRuntimeConfig();
  
  await $fetch(
    `${config.public.apiBase}/establishments/${id}`,
    {
      method: "DELETE",
    }
  );
}
