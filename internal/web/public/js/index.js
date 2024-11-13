let addrsArray = {};
let edit = "0";

loadAddrs();

function createHTML(addr, i) {
    let now = `<i class="bi bi-circle-fill" style="color:var(--bs-gray-500);"></i>`;
    if (addr.Now == 1) {
        now = `<i class="bi bi-check-circle-fill" style="color:var(--bs-success);"></i>`;
    }
    
    let known = '';
    if (addr.Known == 1) {
        known = `checked`;
    }

    // Needs option to use value in js
    let name = `<data id="name${addr.ID}" value="${addr.Name}">${addr.Name}</data>`;
    if (edit == 1) {
        name = `<input id="name${addr.ID}" onchange="editForm(${addr.ID},'')" type="text" class="form-control" value="${addr.Name}">`;
    }

    let html = `
    <tr>
        <td style="opacity: 45%;">${i}.</td>
        <td>${name}</td>
        <td>${addr.Iface}</td>
        <td>
            <a href="http://${addr.IP}" target="_blank">${addr.IP}</a>
        </td>
        <td><a href="/host/${addr.ID}">${addr.Mac}</a></td>
        <td>${addr.Hw}</td>
        <td>${addr.Date}</td>
        <td>${addr.FirstSeen}</td>
        <td>
            <div class="form-check form-switch">
                <input onclick="editForm(${addr.ID}, 'toggle')" class="form-check-input" type="checkbox" ${known}>
            </div>
        </td>
        <td>${now}</td>
    </tr>
    `;
    
    return html;
}

async function loadAddrs() {
    
    const url = '/api/all';
    addrsArray = await (await fetch(url)).json();
    bkpArray = addrsArray;

    field = localStorage.getItem("sortField");
    down = JSON.parse(localStorage.getItem("sortDown"));

    checkNotEmpty(addrsArray);
}

function sortBy(f) {
    field = f;
    if (field != oldField) {
        oldField = field;
        down = !down;
    } else {
        oldField = '';
        down = !down;
    }

    localStorage.setItem("sortDown", down);
    localStorage.setItem("sortField", field);
    checkNotEmpty(addrsArray);
}

function editClick() {

    edit = 1 - edit;
    loadAddrs();
}

function checkNotEmpty(someArray) {

    if (someArray.length > 0) {
        if ((field != null) && (down != null)) {
            sortByAny(someArray);
        } else {
            displayArrayData(someArray);
        }
    }
}

async function editForm(id, known) {
    
    const name = document.getElementById("name"+id).value;
    const url = '/api/edit/'+id+'/'+name+'/'+known;

    // console.log(url);

    await fetch(url);
}
