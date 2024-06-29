var addr = '';
var stop = false;
var portMap = {};
var portArray = [];

// Run, when page loadad
window.addEventListener('load', function() {
    loadSavedPorts();
});

function delRow(id) {
    document.getElementById(id).innerHTML = "";
}

function stopScan() {
    stop = true;
}

async function scanAddr() {
    let begin = document.getElementById("begin").value;
    let end = document.getElementById("end").value;
    let savedPorts = [];

    if (begin == "") {
        begin = 1
    }
    if (end == "") {
        end = 65535
    }
    let port = {};
    stop = false;

    if (portMap != null) {
        savedPorts = Object.keys(portMap);
    }
    // console.log("Saved ports:", savedPorts);

    document.getElementById('stopBtn').style.visibility = "visible";

    for (let i = begin ; i <= end; i++) {

        if (stop) {
            break;
        }

        let url = '/api/port/'+addr+'/'+i;
        port = await (await fetch(url)).json();

        document.getElementById("curPort").innerHTML = "Scanning port "+i;

        if ((port.State) && (!savedPorts.includes(port.Port.toString()))) {
            html = createHTML(port, '', true);
            document.getElementById('tBody').insertAdjacentHTML('afterbegin', html);
        }
    }

    document.getElementById('stopBtn').style.visibility = "hidden";
}

function createHTML(port, i, found) {
    let state = ``;
    let checked = ``;
    let sup = ``;

    if (found) {
        sup = `&nbsp;<sup style="background-color:var(--bs-danger);">new</sup>`;
    }

    if (port.Watch) {
        checked = `checked`;

        if (port.State) {
            state = `<i class="bi bi-check-circle-fill" style="color:var(--bs-success);"></i>`;
        } else {
            state = `<i class="bi bi-dash-circle-fill" style="color:var(--bs-danger);"></i>`;
        }
    } else {
        state = `<i class="bi bi-circle-fill" style="color:var(--bs-gray-500);"></i>`;
    }

    let html = `
    <tr id="row${port.Port}">
        <td style="opacity: 45%;">${i}.</td>
        <td>
            <input name="name" type="text" class="form-control" value="${port.Name}">
        </td>
        <td>
            <a href="http://${addr}:${port.Port}">${port.Port}</a>${sup}
            <input name="port" type="hidden" value="${port.Port}">
            <input name="state" type="hidden" value="${port.State}">
        </td>
        <td>
            ${state}
        </td>
        <td>
            <div class="form-check form-switch">
                <input class="form-check-input" type="checkbox" value="${port.Port}" name="watch" ${checked}>
            </div>
        </td>
        <td>
            <a href="#" onclick="delRow('row${port.Port}')">
                <i class="bi bi-x-circle"></i>
            </a>
        </td>
    </tr>`;
    
    return html;
}

async function loadSavedPorts() {

    addr = document.getElementById("pageAddr").value;
    
    let url = '/api/port/'+addr;
    portMap = await (await fetch(url)).json();
    if (portMap != null) {
        portArray = Object.values(portMap);
    }

    displayArrayData(portArray); // sort.js
}

function sortBy(field) {
    sortByAny(portArray, field); // sort.js
}