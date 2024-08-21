var addrsArray = {};

loadAddrs();

function createHTML(addr, i) {
    let now = '';

    if (addr.Now == 0) {
        now = `<i class="bi bi-circle-fill" style="color:var(--bs-gray-500);"></i>`;
    } else {
        now = `<i class="bi bi-check-circle-fill" style="color:var(--bs-success);"></i>`;
    }
    
    let known = '';

    if (addr.Known == 1) {
        known = `<button type="button" class="btn btn-success">Yes</button>`;
    } else {
        known = `<button type="button" class="btn btn-warning" disabled>No</button>`;
    }

    let html = `
    <tr>
        <td style="opacity: 45%;">${i}.</td>
        <td><input name="name" type="text" class="form-control" value="${addr.Name}"></td>
        <td>${addr.Iface}</td>
        <td>
            <a href="http://${addr.IP}">${addr.IP}</a>
        </td>
        <td>${addr.Mac}</td>
        <td>${addr.Hw}</td>
        <td>${addr.Date}</td>
        <td>${known}</td>
        <td>${now}</td>
    </tr>
    `;
    
    return html;
}

async function loadAddrs() {
    
    let url = '/api/all';
    let addrsMap = await (await fetch(url)).json();
    if (addrsMap != null) {
        addrsArray = Object.values(addrsMap);
    }

    displayArrayData(addrsArray);
}

function sortBy(field) {
    sortByAny(addrsArray, field);
}