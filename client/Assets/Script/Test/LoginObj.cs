using UnityEngine;
using System.Collections;

public class LoginObj : MonoBehaviour {

    public string name_;
    public string passwd_;

	// Use this for initialization
	void Start () {
	
	}
	
	// Update is called once per frame
	void Update () {
	
	}

    public void Init(string name, string passwd)
    {
        name_ = name;
        passwd_ = passwd;
    }

    void OnMouseDown()
    {
        TestManager.Instance.currLoginObj_ = this;
        NetSend.Instance.SendLoginAsk(name_, passwd_);
    }

    public void Destroy()
    {
        Destroy(gameObject);
    }
}
