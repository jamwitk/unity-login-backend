using UnityEngine;

namespace Data
{
    [CreateAssetMenu(fileName = "LoginData", menuName = "Data/LoginData")]
    public class LoginData : ScriptableObject
    {
        public string url;
        public string port;
    }
}
