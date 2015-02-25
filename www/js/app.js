$(document).ready(function() {
    var users = '[{"Id":1,"ProfileUrl":"http://zadane.pl/profil/adassow-6351495","Name":"Adam Sowiński","Nick":""},{"Id":2,"ProfileUrl":"http://znanija.com/profil/Andrea-3","Name":"Andrij Prunko","Nick":""},{"Id":3,"ProfileUrl":"http://zadane.pl/profil/annaef-6083633","Name":"Anna Foltman","Nick":""},{"Id":4,"ProfileUrl":"http://brainly.com.br/perfil/Anna-9","Name":"Anna Skwarek","Nick":""},{"Id":5,"ProfileUrl":"http://zadane.pl/profil/TurboLaczki-6390405","Name":"Arkadiusz Gac","Nick":""},{"Id":6,"ProfileUrl":"http://zadane.pl/profil/DamianGR-5368803","Name":"Damian Gruba","Nick":""},{"Id":7,"ProfileUrl":"http://zadane.pl/profil/Rangoo-6550122","Name":"Dawid Rusnak","Nick":""},{"Id":8,"ProfileUrl":"http://brainly.co.id/profil/dimastoro-6207","Name":"Dimas Widiantoro","Nick":""},{"Id":9,"ProfileUrl":"http://zadane.pl/profil/grzegorzostrow-6390443","Name":"Grzegorz Ostrowski","Nick":""},{"Id":10,"ProfileUrl":"http://zadane.pl/profil/panikowa-6467657","Name":"Iga Kowalska","Nick":""},{"Id":11,"ProfileUrl":"http://zadane.pl/profil/icefyre-6268071","Name":"Jakub Birecki","Nick":""},{"Id":12,"ProfileUrl":"http://brainly.in/profile/Qba-184","Name":"Jakub Bujko","Nick":""},{"Id":13,"ProfileUrl":"http://brainly.com/profile/JPiv-110","Name":"Jakub Piwnik","Nick":""},{"Id":14,"ProfileUrl":"http://zadane.pl/profil/Smajla-5036235","Name":"Joanna Walusiak","Nick":""},{"Id":15,"ProfileUrl":"http://zadane.pl/profil/Roma-134183","Name":"Jolanta Turkiewicz","Nick":""},{"Id":16,"ProfileUrl":"http://zadane.pl/profil/arvaladze-5946337","Name":"Kamil Łukasz","Nick":""},{"Id":17,"ProfileUrl":"http://brainly.com/profile/kamil-5","Name":"Kamil Skruch","Nick":""},{"Id":18,"ProfileUrl":"http://brainly.com/profile/Brooklyn14-130847","Name":"Katie Sagan","Nick":""},{"Id":19,"ProfileUrl":"http://zadane.pl/profil/Amrod-6328910","Name":"Krzysztof Dobrzyński","Nick":""},{"Id":20,"ProfileUrl":"http://zadane.pl/profil/sokpomaranczowy-6387149","Name":"Krzysztof Woliński","Nick":""},{"Id":21,"ProfileUrl":"http://misdeberes.es/perfil/lilismtz-824422","Name":"Liliana Martínez ","Nick":""},{"Id":22,"ProfileUrl":"http://brainly.com.br/perfil/LiviaBR-540411","Name":"Lívia Farias","Nick":""},{"Id":23,"ProfileUrl":"http://zadane.pl/profil/LukaszH-2","Name":"Łukasz Haluch","Nick":""},{"Id":24,"ProfileUrl":"http://zadane.pl/profil/k3nn7-5034172","Name":"Łukasz Lalik","Nick":""},{"Id":25,"ProfileUrl":"http://zadane.pl/profil/lukaszwiewioro-6216358","Name":"Łukasz Wiewiórowski","Nick":""},{"Id":26,"ProfileUrl":"http://brainly.com/profile/mgnat-92963","Name":"Marcin Gnat","Nick":""},{"Id":27,"ProfileUrl":"http://zadane.pl/profil/scanarch-6390066","Name":"Marcin Robaczyński","Nick":""},{"Id":28,"ProfileUrl":"http://zadane.pl/profil/aseroth-5036391","Name":"Marcin Żołna","Nick":""},{"Id":29,"ProfileUrl":"http://zadane.pl/profil/EverydayUser-6060899","Name":"Maria Machlowska","Nick":""},{"Id":30,"ProfileUrl":"http://zadane.pl/profil/Sprzatacz-6385436","Name":"Marta Ryłko","Nick":""},{"Id":31,"ProfileUrl":"http://zadane.pl/profil/KIT-8","Name":"Mateusz Burdzel","Nick":""},{"Id":32,"ProfileUrl":"http://zadane.pl/profil/mycha8-5391369","Name":"Mateusz Zimowski","Nick":""},{"Id":33,"ProfileUrl":"http://nosdevoirs.fr/profil/MathieuM-100756","Name":"Mathieu Milon","Nick":""},{"Id":34,"ProfileUrl":"http://eodev.com/profil/blkmrc-2","Name":"Meriç Balak","Nick":""},{"Id":35,"ProfileUrl":"http://zadane.pl/profil/Michfra-6084117","Name":"Michał  Frasiński","Nick":""},{"Id":36,"ProfileUrl":"http://brainly.com/profile/mackiemesser-44854","Name":"Michał  Łabędź","Nick":""},{"Id":37,"ProfileUrl":"http://zadane.pl/profil/Borek-7","Name":"Michał Borkowski","Nick":""},{"Id":38,"ProfileUrl":"http://zadane.pl/profil/typowymietek-6388724","Name":"Michał Pałka","Nick":""},{"Id":39,"ProfileUrl":"http://zadane.pl/profil/mikemsu-6300088","Name":"Michał Smulski","Nick":""},{"Id":40,"ProfileUrl":"http://zadane.pl/profil/vergilius-5405253","Name":"Michał Żurek","Nick":""},{"Id":41,"ProfileUrl":"http://brainly.ph/profile/Monika-184","Name":"Monika Ambrozowicz","Nick":""},{"Id":42,"ProfileUrl":"http://zadane.pl/profil/mon29ikq-5497016","Name":"Monika Karpierz","Nick":""},{"Id":43,"ProfileUrl":"http://brainly.com/profile/aleksandrawisni-142648","Name":"Ola Wiśniewska","Nick":""},{"Id":44,"ProfileUrl":"http://zadane.pl/profil/oczarnecka-5890944","Name":"Olga Czarnecka","Nick":""},{"Id":45,"ProfileUrl":"http://zadane.pl/profil/patrycyoo-6500703","Name":"Patrycja Szostakowska","Nick":""},{"Id":46,"ProfileUrl":"http://zadane.pl/profil/pmastalerz-6500644","Name":"Paweł Mastalerz","Nick":""},{"Id":47,"ProfileUrl":"http://zadane.pl/profil/pieter1889-5821909","Name":"Piotr Pyznarski","Nick":""},{"Id":48,"ProfileUrl":"http://zadane.pl/profil/przemyslawobor1-6254810","Name":"Przemysław Oborski","Nick":""},{"Id":49,"ProfileUrl":"http://brainly.com/profile/Scryt-32498","Name":"Przemysław Toś","Nick":""},{"Id":50,"ProfileUrl":"http://zadane.pl/profil/cieciwka4-6388720","Name":"Radek Cięciwa","Nick":""},{"Id":51,"ProfileUrl":"http://zadane.pl/profil/betterthanperfect-6500625","Name":"Roman Barzyczak","Nick":""},{"Id":52,"ProfileUrl":"http://zadane.pl/profil/TypowySebaxD-6388721","Name":"Sebastian Czoch","Nick":""},{"Id":53,"ProfileUrl":"http://zadane.pl/profil/drosth-6348157","Name":"Tomasz Drozd","Nick":""},{"Id":54,"ProfileUrl":"http://zadane.pl/profil/tomaszek18-6387947","Name":"Tomasz Kraus","Nick":""},{"Id":55,"ProfileUrl":"http://zadane.pl/profil/winglot-5366969","Name":"Wojciech Inglot","Nick":""},{"Id":56,"ProfileUrl":"http://brainly.com/profile/voytek-38246","Name":"Wojtek Skalski","Nick":""},{"Id":57,"ProfileUrl":"http://zadane.pl/profil/pannandrzej-6449039","Name":"Dawid Gaweł","Nick":""},{"Id":58,"ProfileUrl":"http://brainly.com/profile/jasongreen1-124588","Name":"Jason Green","Nick":""}]';
    var scores = '[{"UserId":34,"Answers7Days":46,"Answers30Days":55,"Answers90Days":100},{"UserId":30,"Answers7Days":25,"Answers30Days":47,"Answers90Days":49},{"UserId":20,"Answers7Days":22,"Answers30Days":91,"Answers90Days":191},{"UserId":9,"Answers7Days":19,"Answers30Days":124,"Answers90Days":124},{"UserId":51,"Answers7Days":12,"Answers30Days":23,"Answers90Days":89},{"UserId":8,"Answers7Days":9,"Answers30Days":47,"Answers90Days":208},{"UserId":26,"Answers7Days":8,"Answers30Days":26,"Answers90Days":98},{"UserId":29,"Answers7Days":6,"Answers30Days":6,"Answers90Days":10},{"UserId":33,"Answers7Days":5,"Answers30Days":7,"Answers90Days":15},{"UserId":15,"Answers7Days":2,"Answers30Days":9,"Answers90Days":18},{"UserId":57,"Answers7Days":2,"Answers30Days":2,"Answers90Days":42},{"UserId":21,"Answers7Days":1,"Answers30Days":3,"Answers90Days":14},{"UserId":1,"Answers7Days":0,"Answers30Days":22,"Answers90Days":26},{"UserId":16,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":17,"Answers7Days":0,"Answers30Days":19,"Answers90Days":20},{"UserId":18,"Answers7Days":0,"Answers30Days":0,"Answers90Days":1},{"UserId":19,"Answers7Days":0,"Answers30Days":1,"Answers90Days":13},{"UserId":22,"Answers7Days":0,"Answers30Days":1,"Answers90Days":1},{"UserId":24,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":25,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":27,"Answers7Days":0,"Answers30Days":0,"Answers90Days":46},{"UserId":28,"Answers7Days":0,"Answers30Days":0,"Answers90Days":1},{"UserId":31,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":32,"Answers7Days":0,"Answers30Days":0,"Answers90Days":2},{"UserId":35,"Answers7Days":0,"Answers30Days":0,"Answers90Days":7},{"UserId":36,"Answers7Days":0,"Answers30Days":0,"Answers90Days":2},{"UserId":37,"Answers7Days":0,"Answers30Days":0,"Answers90Days":15},{"UserId":38,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":39,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":40,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":41,"Answers7Days":0,"Answers30Days":1,"Answers90Days":3},{"UserId":42,"Answers7Days":0,"Answers30Days":0,"Answers90Days":16},{"UserId":43,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":44,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":45,"Answers7Days":0,"Answers30Days":11,"Answers90Days":12},{"UserId":46,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":47,"Answers7Days":0,"Answers30Days":3,"Answers90Days":5},{"UserId":48,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":49,"Answers7Days":0,"Answers30Days":2,"Answers90Days":4},{"UserId":50,"Answers7Days":0,"Answers30Days":0,"Answers90Days":10},{"UserId":52,"Answers7Days":0,"Answers30Days":0,"Answers90Days":1},{"UserId":53,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":54,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":55,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":56,"Answers7Days":0,"Answers30Days":1,"Answers90Days":2},{"UserId":58,"Answers7Days":0,"Answers30Days":25,"Answers90Days":26},{"UserId":23,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":2,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":3,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":4,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":5,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":6,"Answers7Days":0,"Answers30Days":4,"Answers90Days":30},{"UserId":7,"Answers7Days":0,"Answers30Days":1,"Answers90Days":56},{"UserId":10,"Answers7Days":0,"Answers30Days":0,"Answers90Days":0},{"UserId":11,"Answers7Days":0,"Answers30Days":0,"Answers90Days":10},{"UserId":12,"Answers7Days":0,"Answers30Days":1,"Answers90Days":2},{"UserId":13,"Answers7Days":0,"Answers30Days":4,"Answers90Days":28},{"UserId":14,"Answers7Days":0,"Answers30Days":1,"Answers90Days":1}]';

    var winnerTemplate = getWinnerTemplate();
    var looserTemplate = getLooserTemplate();

    // $.get("http://188.166.53.60:8080/users", function(users) {
    //     $.get("http://188.166.53.60:8080/scores", function(scores) {
            var usersById = new Array;

            users = JSON.parse(users);
            scores = JSON.parse(scores);

            for (i in users) {
                usersById[users[i].Id] = users[i]
            }

            for (i in scores) {
                var rowData = new Array;
                score = scores[i]
                user = usersById[score["UserId"]]

                rowData[0] = parseInt(i) + 1;
                rowData[1] = user["Nick"]
                rowData[2] = user["Name"]
                rowData[3] = score["Answers7Days"]
                rowData[4] = score["Answers30Days"]
                rowData[5] = score["Answers90Days"]

                if (i < 3) {
                    var row = winnerTemplate({
                       place: parseInt(i) + 1,
                       avatar_url: "http://zadane.pl/img/avatars/100-ONA.png",
                       nick: user['Nick'],
                       name: user['Name'],
                       answers7days: score['Answers7Days'],
                       answers30days: score['Answers30Days'],
                       answers90days: score['Answers90Days']
                    });
                    $(".winners-table > tbody:last").append(row);
                } else {
                    var row = looserTemplate({
                       place: parseInt(i) + 1,
                       avatar_url: "http://zadane.pl/img/avatars/100-ONA.png",
                       nick: user['Nick'],
                       name: user['Name'],
                       answers7days: score['Answers7Days'],
                       answers30days: score['Answers30Days'],
                       answers90days: score['Answers90Days']
                    });
                    $(".loosers-table > tbody:last").append(row);
                }

            }
    //     });
    // });

    function rowFromRowData(place, rowData) {
        var row = "<tr>";

        for (i in rowData) {
            row = row + "<td>";
            row = row + rowData[i];
            row = row + "</td>";
        }

        row = row + "</tr>"

        return row;
    }

    function getWinnerTemplate() {
        return _.template('<tr>' +
                   '<td class="place"><%= place %></td>' +
                   '<td class="avatar">' +
                       '<div class="avatar-image">' +
                           '<img src="<%= avatar_url %>" class="img-circle avatar-img" />' +
                           '<img src="img/icon_prize_0<%= place %>.png" class="prize-img" />' +
                       '</div>' +
                       '<div class="user-name">' +
                           '<p class="nick"><%= nick %></p>' +
                           '<p class="name"><%= name %></p>' +
                       '</div>' +
                   '</td>' +
                   '<td class="score"><%= answers7days %></td>' +
                   '<td class="score"><%= answers30days %></td>' +
                   '<td class="score"><%= answers90days %></td>' +
               '</tr>');
    }

    function getLooserTemplate() {
        return _.template('<tr>' +
                    '<td class="place"><%= place %></td>' +
                    '<td class="avatar">' +
                        '<img src="<%= avatar_url %>" class="img-circle" />' +
                        '<div class="user-name">' +
                            '<p class="nick"><%= nick %></p>' +
                            '<p class="name"><%= name %></p>' +
                        '</div>' +
                    '</td>' +
                    '<td class="score"><%= answers7days %></td>' +
                    '<td class="score"><%= answers30days %></td>' +
                    '<td class="score"><%= answers90days %></td>' +
                '</tr>');
    }


   // var compiled = _.template(winnerTemplate);

   // console.log(compiled({
   //     place: 1,
   //     avatar_url: "http://example.com",
   //     nick: "k3nn7",
   //     name: "Łukasz Lalik",
   //     answers7days: 30,
   //     answers30days: 90,
   //     answers90days: 110
   // }));
});
