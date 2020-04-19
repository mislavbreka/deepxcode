$(document).ready(function() {
    
    $("#menu").click(function(){
        if ($('#navcontent').hasClass('hide')){
            $('#navcontent').removeClass('hide');
            $('#menu').addClass('change');
        }else{
            $('#navcontent').addClass('hide');
            $('#menu').removeClass('change');
        }
    })
    $("#search").click(function(){
        if ($('#navsearch').hasClass('hide')){
            $('#navsearch').removeClass('hide');
            $('#navresults').removeClass('hide');
            $('#search').addClass('change');7
        }else{
            $('#navsearch').addClass('hide');
            $('#navresults').addClass('hide');
            $('#search').removeClass('change');
        }
    })

    $("#login-icon").click(function(){
        window.location.replace("/login/")
    })
    $("#register-icon").click(function(){
        window.location.replace("/register/")
    })
    $("#profile-icon").click(function(){
        window.location.replace("/profile/")
    })
    $("#data-icon").click(function(){
        window.location.replace("/data/")
    })
    $("#github-icon").click(function(){
        window.location.href = "https://github.com/mislavbreka"
    })
    $("#atom-logo").click(function(){
        window.location.replace("/")
    })
    
    $('#searchinput').on('keyup', function(e) {
        e.preventDefault();

        var input = document.getElementById('searchinput');
        var filter = input.value.toUpperCase();
        var results = document.getElementById("navresults");
        var li = results.getElementsByTagName('li');

        var display = 3
        // Loop through all list items, and display those who don't match the search query
        for (i = 0; i < li.length; i++) {
            a = li[i].getElementsByTagName("a")[0];
            txtValue = a.textContent || a.innerText;
            if (txtValue.toUpperCase().indexOf(filter) > -1) {
                display--;
                li[i].style.display = "grid";
            } else {
                li[i].style.display = "none";
            }
            if (display<0) {
                break
            }
        }

    });
});