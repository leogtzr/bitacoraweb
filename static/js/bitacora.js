$(document).ready(function () {

    $('#alert').hide();
    $('#alert_error').hide();

    if (existsEmployeeParamInURI()) {
        console.log('Exists employee query string ... ');
        const urlParams = new URLSearchParams(window.location.search);
        
        var empID = urlParams.get('employee');

        $.ajax({
            url: 'http://localhost:8083/api/entries/' + empID,
            type: 'GET',
            data: {},
            success: function(data) {
                console.log(data);

                var entriesTable = $('#entries_table');

                var entries = data.entries;
                for (var i = 0; i < entries.length; i++) {
                    var entry = entries[i];
                    var whoRow = '<tr><td><b>' + data.name + '</b></td>';
                    whoRow += '<td>' + entry.description + '</td>';
                    whoRow += '<td>' + entry.date + '</td></tr>';
                                                                                                                                
                    entriesTable.append(whoRow);
                }
            },
            error: function(data) {
                console.log('woops! :( alrpvw' + data);
            }
        });

    }

    if ($('#date').length) {
        $('#date').datepicker({
            format: "yyyy/mm/dd",
            weekStart: 1,
            todayBtn: "linked",
            todayHighlight: true
        });
    }

    $(".today").click();

    var persons = $('#persons');
    if (persons.length) {
        $.ajax({
            url: 'http://localhost:8083/api/employees',
            type: 'GET',
            data: {},
            success: function(data) {
                for (i = 0; i < data.length; i++) {
                    $('#persons').append($('<option name="' + data[i]._id + '">').append(data[i].name));
                }
            },
            error: function(data) {
                console.log(data);
                console.log('woops! :(' + data);
            }
        });
    }

    
    var teamEntries = $('#team');
    if (teamEntries.length) {
        console.log('Exists ... emp id');
        $.ajax({
            url: 'http://localhost:8083/api/employees',
            type: 'GET',
            data: {},
            success: function(data) {
                var types = data;
                for (i = 0; i < types.length; i++) {
                    var comment = types[i].name;
                    comment = comment.substring(0, 20) + '...';
                    $('#team').append(
                        $('<li name="1" class="list-group-item">').append(
                            $('<a>').attr('href','/entries.html?employee=' + types[i]._id).append(
                                $('<h5>').attr('class', '').append(types[i].name)
                    )));
                }
            },
            error: function(data) {
                console.log('woops! :(');
                console.log(data);
            }
        });
    }


    $('#addentry').on('submit', function(e) {

        e.preventDefault();
        var text = $('#interactiontext').val();
        var empId = $('#persons').find(":selected").attr('name');
        var date = $('#date').val();

        $.ajax({
            url: 'http://localhost:8083/api/entrytoemp',
            type: 'POST',
            data: {id: empId, when: date, description: text, notes: []},
            success: function(data) {
                $('#interactiontext').val('');
            },
            error: function(data) {
                console.log(data);
                $("#alert_error").fadeTo(2000, 500).slideUp(500, function() {
                    $("#alert_error").slideUp(500);
                });
            }
        });

    });

});

function existsEmployeeParamInURI() {
    var url = new URL(window.location.href);
    if (url.searchParams.get('employee')) {
        return true;
    } else {
        return false;
    }
}
