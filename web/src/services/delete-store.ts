export async function deleteStore(id: string): Promise<void> {
  const config = useRuntimeConfig();
  
  await $fetch(
    `${config.public.apiBase}/stores/${id}`,
    {
      method: "DELETE",
    }
  );
}
