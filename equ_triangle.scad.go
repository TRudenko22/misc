// Equilateral Triangle

module equ_triangle(side_length,corner_radius,height){
    translate([5,corner_radius]){
         hull(){
           cylinder(r=corner_radius,h=height);
           rotate([0,0,60])translate([side_length-corner_radius*2,0,0])cylinder(r=corner_radius,h=height);   
           rotate([0,0,120])translate([side_length-corner_radius*2,0,0])cylinder(r=corner_radius,h=height);
      }
   }
}

