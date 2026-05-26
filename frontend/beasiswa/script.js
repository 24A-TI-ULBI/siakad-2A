fetch("http://127.0.0.1:8080/api/beasiswa")
.then(response => response.json())
.then(data => {

  const container = document.getElementById("beasiswa-list");

  data.forEach(item => {

    container.innerHTML += `
      <div class="card">

        <h3>${item.nama}</h3>

        <p>
          <strong>Syarat:</strong>
          ${item.syarat}
        </p>

        <p>
          <strong>Deadline:</strong>
          ${item.deadline}
        </p>

      </div>
    `;

  });

})
.catch(error => console.log(error));


function daftarBeasiswa(){

  const nama = document.getElementById("nama").value;

  const npm = document.getElementById("npm").value;

  const pilihan = document.getElementById("pilihan").value;

  const status = document.getElementById("status-message");

  if(nama === '' || npm === ''){

    status.innerHTML = `
      <div class="success-message">
        Please complete all fields.
      </div>
    `;

    return;
  }

  status.innerHTML = `
    <div class="success-message">
      Registration submitted successfully.
    </div>
  `;
}

function hapusFile(){

  document.getElementById("motivasi").value = "";

}

async function daftarBeasiswa(){

  const data = {

    nama:
      document.getElementById("nama").value,

    npm:
      document.getElementById("npm").value,

    email:
      document.getElementById("email").value,

    semester:
      document.getElementById("semester").value,

    prodi:
      document.getElementById("prodi").value,

    ipk:
      document.getElementById("ipk").value,

    beasiswa:
      document.getElementById("pilihan").value
  };

  const response = await fetch(
    "http://127.0.0.1:8080/beasiswa/daftar",
    {
      method:"POST",

      headers:{
        "Content-Type":"application/json"
      },

      body:JSON.stringify(data)
    }
  );

  const result = await response.json();
  console.log(result);
  alert(result.message || "Pendaftaran berhasil");
}

async function cekStatus(){

  const npm =
    document.getElementById("cekNpm").value;

  const response = await fetch(
    `http://127.0.0.1:8080/beasiswa/status/${npm}`
  );

  const result = await response.json();

  const hasil =
    document.getElementById("hasilStatus");

  if(result.message){

    hasil.innerHTML = `
      <p>${result.message}</p>
    `;

    return;
  }

  hasil.innerHTML = `
    <div class="status-result">

      <h3>${result.data.nama}</h3>

      <p><strong>NPM:</strong> ${result.data.npm}</p>

      <p><strong>Scholarship:</strong> ${result.data.beasiswa}</p>

      <p><strong>Status:</strong> ${result.data.status}</p>
          </div>
  `;
}

async function updateStatus() {

  const npm =
    document.getElementById("updateNpm").value;

  const status =
    document.getElementById("updateStatus").value;

  try {

    const response = await fetch(
      `http://127.0.0.1:8080/beasiswa/status/${npm}`,
      {
        method: "PUT",

        headers: {
          "Content-Type": "application/json"
        },

        body: JSON.stringify({
          status: status
        })
      }
    );

    const result = await response.json();

    console.log(result);

    alert(result.message || "Status berhasil diupdate");

  } catch (error) {

    console.log(error);

    alert("Gagal update status");

  }

}

async function deletePendaftaran() {

  const npm =
    document.getElementById("deleteNpm").value;

  try {

    const response = await fetch(
      `http://127.0.0.1:8080/beasiswa/${npm}`,
      {
        method: "DELETE"
      }
    );

    const result = await response.json();

    console.log(result);

    alert("Data berhasil dihapus");

  } catch (error) {

    console.log(error);

    alert("Gagal menghapus data");

  }

}