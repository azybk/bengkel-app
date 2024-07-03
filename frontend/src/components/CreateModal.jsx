const CreateModal = () => {
  return (
    <template>
      <div class="absolute top-0 left-0 w-full h-screen flex items-center justify-center">
        <div class="absolute top-0 left-0 w-full h-screen bg-gray-900 opacity-10 z-10"></div>
        <div class="max-w-sm bg-white shadow-xl rounded overflow-hidden flex-1 z-20 border border-gray-300">
          <div class="">
            <div class="py-3 px-2 bg-gray-100">
              <p class="text-sm text-gray-800">Tambah Historical</p>
            </div>
            <div class="px-2 py-4">
              <div class="mb-2">
                <p class="text-xs">Customer</p>
                <select class="py-1.5 px-3 text-sm border border-gray-300 w-full rounded">
                  <option>Sefna</option>
                </select>
              </div>
              <div class="mb-2">
                <p class="text-xs">PIC</p>
                <input
                  type="text"
                  class="py-1 px-3 text-sm border border-gray-300 w-full rounded"
                />
              </div>
              <div>
                <p class="text-xs">Plat Nomor</p>
                <input
                  type="text"
                  class="py-1 px-3 text-sm border border-gray-300 w-full rounded"
                />
              </div>
              <div>
                <p class="text-xs">Catatan</p>
                <textarea class="py-1 px-3 text-sm border border-gray-300 w-full rounded"></textarea>
              </div>
            </div>
          </div>
          <div class="grid grid-cols-2 py-2 px-2 bg-gray-100 gap-2">
            <button class="border border-gray-300 text-gray-600 bg-white text-sm py-1 rounded hover:shadow">
              Batal
            </button>
            <button class="border border-gray-300 text-white bg-blue-400 text-sm py-1 rounded hover:shadow">
              Simpan
            </button>
          </div>
        </div>
      </div>
    </template>
  );
};

export default CreateModal;
