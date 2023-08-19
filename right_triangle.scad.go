// Right Triangle

module right_triangle(side_1, side_2, corner_radius, height){
    translate([corner_radius, corner_radius, 0]){
        hull(){
            cylinder(r=corner_radius,h=height);
            translate([side_1-corner_radius*2,0,0])cylinder(r=corner_radius,h=height); 
            translate([0,side_1-corner_radius*2,0])cylinder(r=corner_radius,h=height);
        }
    }
}
