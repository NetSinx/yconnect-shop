import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { User } from 'src/app/interfaces/user';
import { UserService } from 'src/app/services/user/user.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  user_id: string | null
  user: User | undefined

  constructor(private userService: UserService, private route: ActivatedRoute) {
    this.user_id = this.route.snapshot.paramMap.get("userId")
  }

  ngOnInit(): void {
    this.getUser(this.user_id!)
  }

  public getUser(user_id: string): void {
    this.userService.getUser(user_id).subscribe(
      resp => this.user = resp.data
    )
  }
}
