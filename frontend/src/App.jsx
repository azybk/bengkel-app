import "./App.css";

const App = () => {
  return (
    <main>
      <div class="max-w-7xl mx-auto shadow p-4 rounded-xl my-2 bg-white">
        <div class="mb-4">
          <p>Bengkel Uchook</p>
          <p class="text-sm text-gray-500">
            System informasi perbengkelan Uchok
          </p>
        </div>
        <div class="max-w-sm flex mb-4">
          <input
            type="text"
            class="py-1.5 px-2 border border-gray-300 rounded-l text-sm flex-1"
            placeholder="Masukkan nomor kendaraan"
          />
          <button class="py-1.5 px-2 bg-blue-400 text-white text-sm rounded-r">
            Cari kendaraan
          </button>
        </div>
        <div class="mb-4 max-w-xl">
          <table class="text-sm w-full">
            <tr>
              <td class="text-gray-600 w-56">No Kendaraan</td>
              <td class="w-3">:</td>
              <td>
                <input
                  type="text"
                  class="py-1.5 px-2 border border-gray-300 rounded text-sm flex-1 w-full"
                  placeholder="Masukkan nomor kendaraan"
                />
              </td>
            </tr>
            <tr>
              <td class="text-gray-600 w-56">Brand</td>
              <td class="w-3">:</td>
              <td>
                <input
                  type="text"
                  class="py-1.5 px-2 border border-gray-300 rounded text-sm flex-1 w-full"
                  placeholder="Brand"
                />
              </td>
            </tr>
          </table>
        </div>
        <div>
          <table class="text-sm w-full">
            <tr class="bg-gray-100">
              <td class="border border-gray-300 py-1 px-2">Tanggal</td>
              <td class="border border-gray-300 py-1 px-2">Plat Nomor</td>
              <td class="border border-gray-300 py-1 px-2">Customer</td>
              <td class="border border-gray-300 py-1 px-2">Notes</td>
              <td class="border border-gray-300 py-1 px-2">PIC</td>
            </tr>
            <tr>
              <td class="border border-gray-300 py-1 px-2">17-12-2020</td>
              <td class="border border-gray-300 py-1 px-2">B 123 ABC</td>
              <td class="border border-gray-300 py-1 px-2">Sefna</td>
              <td class="border border-gray-300 py-1 px-2">
                Penggantian sparepart
              </td>
              <td class="border border-gray-300 py-1 px-2">Suryono</td>
            </tr>
          </table>
          <div class="flex justify-start mt-1">
            <button class="py-1.5 px-2 bg-gray-100 text-gray-600 border border-gray-300 text-sm rounded hover:shadow">
              Tambah histori
            </button>
          </div>
        </div>
      </div>
    </main>
  );
};

export default App;
